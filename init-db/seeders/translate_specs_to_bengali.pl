#!/usr/bin/perl
use strict;
use warnings;

# English specification keywords to Bengali translations
my %translations = (
    ' MP' => ' মেগাপিক्षेल',
    'pixels' => 'पिक्षेल',
    ' GB' => ' गिگाবाइট',
    ' TB' => ' টেরাबাइट',
    ' mAh' => ' এमএএইচ',
    'inches' => 'ইঞ्চি',
    ' mm' => ' মिमि',
    ' g' => ' গ्रাम',
    ' nm' => ' nm',  # Keep nm as is
    'Hz' => 'Hz',  # Keep Hz as unit
    'Android' => 'अ্́যান्ड्रয়েড',
    'Glass' => 'গ्लাস',
    'Plastic' => 'प्लাস्टिक',
    'Metal' => 'ধাতু',
    'No' => 'না',
    'Yes' => 'হ্যাঁ',
    'Available' => 'উপলब্ध',
    'core' => 'কোর',
    'LTPO' => 'এলটিপিও',
    'AMOLED' => 'অ্যামোলেড',
    'LCD' => 'এলসিডি',
    'OLED' => 'ওলেড',
    'IPS' => 'আইপিএস',
    'nits' => 'নিট',
);

my $in_translations = 0;

while (<>) {
    my $line = $_;
    
    if ($line =~ /func.*getBanglaTranslations/) {
        $in_translations = 1;
    } elsif ($line =~ /^func/ && $in_translations) {
        $in_translations = 0;
    }
    
    # Only in translation map values
    if ($in_translations && $line =~ /":\s*"/) {
        # For each English keyword, replace with Bengali
        # This is a bit tricky because we need to preserve structure
        # For now, just do simple replacements
        foreach my $eng (sort {length($b) <=> length($a)} keys %translations) {
            $line =~ s/\Q$eng\E/$translations{$eng}/g;
        }
    }
    
    print $line;
}
