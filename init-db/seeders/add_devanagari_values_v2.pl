#!/usr/bin/perl
use strict;
use warnings;

my %a2d = ('0'=>'०','1'=>'१','2'=>'२','3'=>'३','4'=>'४','5'=>'५','6'=>'६','7'=>'७','8'=>'८','9'=>'९');

my $in_translations = 0;

while (<>) {
    my $line = $_;
    
    # Track if we're in getBanglaTranslations function
    if ($line =~ /func.*getBanglaTranslations/) {
        $in_translations = 1;
        print $line;
        next;
    } elsif ($line =~ /^func/ && $in_translations) {
        $in_translations = 0;
    }
    
    # Within getBanglaTranslations: convert ASCII to Devanagari in map entries
    if ($in_translations && $line =~ /":\s*"/) {
        # Match pattern: "key": "value"
        # The pattern is: starts with optional whitespace, then KEY in quotes, then : then VALUE in quotes
        if ($line =~ /^(\s*)"([^"]*)"(\s*:\s*")([^"]*)(".*)$/) {
            my ($indent, $key, $sep, $value, $rest) = ($1, $2, $3, $4, $5);
            # Convert digits in value to Devanagari
            foreach my $digit (keys %a2d) {
                $value =~ s/$digit/$a2d{$digit}/g;
            }
            $line = $indent . '"' . $key . '"' . $sep . $value . $rest;
        }
    }
    
    print $line;
}
