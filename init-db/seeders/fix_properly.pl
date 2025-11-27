#!/usr/bin/perl
use strict;
use warnings;

my %ascii_to_devanagari = (
    '0' => '०', '1' => '१', '2' => '२', '3' => '३', '4' => '४',
    '5' => '५', '6' => '६', '7' => '७', '8' => '८', '9' => '९',
);

my %devanagari_to_ascii = reverse %ascii_to_devanagari;

my $in_bangla = 0;
my $in_specs = 0;

while (<>) {
    my $line = $_;
    
    # Track which map we're in
    if ($line =~ /func.*getBanglaTranslations.*\{/) {
        $in_bangla = 1;
        $in_specs = 0;
    } elsif ($line =~ /specs\s*:=\s*map\[string\]string\s*\{/) {
        $in_bangla = 0;
        $in_specs = 1;
    } elsif ($line =~ /^\s*\}/) {
        $in_bangla = 0;
        $in_specs = 0;
    }
    
    # Within getBanglaTranslations: convert ASCII to Devanagari in VALUES only
    if ($in_bangla && $line =~ /^(\s*"[^"]*":\s*")([^"]*)(".*)$/) {
        my ($before, $value, $after) = ($1, $2, $3);
        # Convert digits in value only
        foreach my $digit (keys %ascii_to_devanagari) {
            $value =~ s/$digit/$ascii_to_devanagari{$digit}/g;
        }
        $line = $before . $value . $after . "\n";
    }
    
    # Within specs map: convert Devanagari to ASCII everywhere  
    if ($in_specs) {
        foreach my $dev (keys %devanagari_to_ascii) {
            $line =~ s/$dev/$devanagari_to_ascii{$dev}/g;
        }
    }
    
    print $line;
}
